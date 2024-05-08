import { Component, OnDestroy, OnInit } from '@angular/core';
import {
  FormControl,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { LruService } from './lru.service';
import { MatTableModule } from '@angular/material/table';
import { CommonModule } from '@angular/common';
import { WebSocketSubject, webSocket } from 'rxjs/webSocket';

@Component({
  selector: 'app-lru',
  standalone: true,
  imports: [
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
    MatTableModule,
    CommonModule,
  ],
  templateUrl: './lru.component.html',
  styleUrl: './lru.component.scss',
})
export class LruComponent implements OnInit, OnDestroy {
  dataForm = new FormGroup({
    key: new FormControl('', Validators.required),
    value: new FormControl('', Validators.required),
  });
  data: DataItem[] = [];
  originalData: DataItem[] = [];
  expiredKeys: string[] = []
  private socket$!: WebSocketSubject<any>;
  constructor(private lruService: LruService,
  ) {}

  ngOnInit(): void {
    this.socket$ = webSocket('ws://localhost:8081/ws');
    this.getAlldata();
    this.socket$.subscribe(
      (message) => { 
        this.getAlldata();
        console.log('Received message:', message) 
      },
      (error) => console.error('WebSocket error:', error),
      () => console.log('WebSocket connection closed.')
    );
  }

  ngOnDestroy() {
    //this.socket.close();
  }

  public getAlldata() {
    this.originalData=[];
    this.data=[];
    // Store a backup copy of the original data
    
     this.lruService.getAllData().subscribe({
      next: (result: any) => {
        this.originalData= result;
        this.data = [...this.originalData];
      },
      error: (error: any) => {
        console.error("Error al comparar las versiones", error);
      }
    }); 
  }

  public add(): void {
    let data: object = {
      key: this.dataForm.get('key')?.value!,
      value: this.dataForm.get('value')?.value!,
    };
    /* let response= this.lruService.sendData(data); */
    this.lruService.sendData(data).subscribe((result) => {
      if(result){
         this.getAlldata();
         this.dataForm.reset()
      }
    });
  }

  public searchByKey(event: Event): void {
    let key = (event.target as HTMLInputElement).value;
    if (!key.trim()) {
      // If the search value is empty, restore the original data
      this.data = [...this.originalData];
      return;
    }
    // Filter the data based on the provided value
    this.data = this.originalData.filter(
      (item) =>
        item['key'].toLowerCase().includes(key.toLowerCase())
    );
  }
}

export interface DataItem {
  key: string;
  value: string;
  expiration?: string;
}
