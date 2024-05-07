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
import {MatSnackBar} from '@angular/material/snack-bar';

@Component({
  selector: 'app-library',
  standalone: true,
  imports: [FormsModule, ReactiveFormsModule, MatFormFieldModule, MatInputModule, MatButtonModule, MatIconModule],
  templateUrl: './library.component.html',
  styleUrl: './library.component.scss'
})
export class LibraryComponent {
  comparisonResult: string = '';
  constructor(private _snackBar: MatSnackBar) {}

  dataForm = new FormGroup({
    version1: new FormControl('', Validators.required),
    version2: new FormControl('', Validators.required),
  });

  onKeyDown(event: KeyboardEvent) {
    // Permitir solo números, punto, retroceso y teclas de flecha
    const allowedChars = /[0-9\.]|Backspace|Tab|ArrowLeft|ArrowRight|ArrowUp|ArrowDown/;
    const key = event.key;
    if (!allowedChars.test(key)) {
      event.preventDefault();
    }
  }

  compareVersions(): void {
    if(this.dataForm.valid){

    }
    //consumir back, almacenar en comparisonResult
  }
}