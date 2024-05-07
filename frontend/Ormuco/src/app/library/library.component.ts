import { Component } from '@angular/core';
import {
  FormControl,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatIconModule} from '@angular/material/icon';
import { LibraryService } from './library.service';

@Component({
  selector: 'app-library',
  standalone: true,
  imports: [FormsModule, ReactiveFormsModule, MatFormFieldModule, MatInputModule, MatButtonModule, MatIconModule],
  templateUrl: './library.component.html',
  styleUrl: './library.component.scss',
})
export class LibraryComponent {
  comparisonResult: string = '';
  constructor(private libraryService: LibraryService
  ) {}

  dataForm = new FormGroup({
    version1: new FormControl('', Validators.required),
    version2: new FormControl('', Validators.required),
  });

  onKeyDown(event: KeyboardEvent) {
    const allowedChars = /[0-9\.]|Backspace|Tab|ArrowLeft|ArrowRight|ArrowUp|ArrowDown/;
    const key = event.key;
    if (!allowedChars.test(key)) {
      event.preventDefault();
    }
    this.comparisonResult="";
  }

  compareVersions(): void {
    const version1 = this.dataForm.get('version1')?.value!;
      const version2 = this.dataForm.get('version2')?.value!;
      this.libraryService.getCompareVersions(version1, version2).subscribe({
        next: (result: any) => {
          this.comparisonResult = result;
        },
        error: (error: any) => {
          console.error("Error al comparar las versiones", error);
        }
      });
  }
}