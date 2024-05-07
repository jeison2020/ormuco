import { Component, ViewChild, ElementRef, AfterViewInit } from '@angular/core';
import {
  FormControl,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import Chart from 'chart.js/auto';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-graph',
  standalone: true,
  imports: [
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
  ],
  templateUrl: './graph.component.html',
  styleUrl: './graph.component.scss',
})
export class GraphComponent implements AfterViewInit {
  constructor(private _snackBar: MatSnackBar) {}

  @ViewChild('lineChart') private lineChartRef!: ElementRef;

  dataForm = new FormGroup({
    x1: new FormControl(0, Validators.required),
    x2: new FormControl(0, Validators.required),
    x3: new FormControl(0, Validators.required),
    x4: new FormControl(0, Validators.required),
  });

  overlap: boolean = false;
  showText: boolean = true;
  private lineChart: Chart | null = null;

  ngAfterViewInit(): void {
    this.updateChart();
    // Subscribe to changes in the form
    this.dataForm.valueChanges.subscribe(() => {
      // Logic to handle changes in the form
      if (this.showText) {
        this.showText = false;
      }
      this.destroyChart();
    });
  }

  destroyChart() {
    if (this.lineChart) {
      this.lineChart.destroy();
    }
  }

  updateChart(): void {
    const xLabels = ['x1', 'x2', 'x3', 'x4'];
    this.destroyChart();

    if (this.dataForm.valid) {
      if (this.dataForm.get('x1')?.value! <= this.dataForm.get('x2')?.value!) {
        if (
          this.dataForm.get('x3')?.value! <= this.dataForm.get('x4')?.value!
        ) {
          this.showText = true;
          // Check if lines overlap
          this.overlap =
            Math.max(
              this.dataForm.get('x1')?.value!,
              this.dataForm.get('x3')?.value!
            ) <=
              Math.min(
                this.dataForm.get('x2')?.value!,
                this.dataForm.get('x4')?.value!
              ) &&
            Math.max(
              this.dataForm.get('x3')?.value!,
              this.dataForm.get('x1')?.value!
            ) <=
              Math.min(
                this.dataForm.get('x4')?.value!,
                this.dataForm.get('x2')?.value!
              );

          const overlapLabel = `Overlap between ${Math.max(
            this.dataForm.get('x1')?.value!,
            this.dataForm.get('x3')?.value!
          )} and ${Math.min(
            this.dataForm.get('x2')?.value!,
            this.dataForm.get('x4')?.value!
          )}`;

          const datasets = [
            {
              label: 'Line 1',
              data: [
                { x: this.dataForm.get('x1')?.value!, y: 0 },
                { x: this.dataForm.get('x2')?.value!, y: 0 },
              ],
              borderColor: '#6F83BF',
              backgroundColor: '#6F83BF',
              pointRadius: 5,
              showLine: true,
              borderWidth: 2,
            },
            {
              label: 'Line 2',
              data: [
                { x: this.dataForm.get('x3')?.value!, y: 0 },
                { x: this.dataForm.get('x4')?.value!, y: 0 },
              ],
              borderColor: '#021E73',
              backgroundColor: '#021E73',
              pointRadius: 6,
              borderWidth: 2,
              showLine: true,
            },
          ];

          // Add the dataset for overlap only if there is overlap
          if (this.overlap) {
            datasets.push({
              label: overlapLabel,
              data: [
                {
                  x: Math.max(
                    this.dataForm.get('x1')?.value!,
                    this.dataForm.get('x3')?.value!
                  ),
                  y: 0,
                },
                {
                  x: Math.min(
                    this.dataForm.get('x2')?.value!,
                    this.dataForm.get('x4')?.value!
                  ),
                  y: 0,
                },
              ],
              borderColor: '#F24957',
              backgroundColor: '#F24957',
              pointRadius: 10,
              showLine: true,
              borderWidth: 5,
            });
          }

          const ctx = this.lineChartRef.nativeElement.getContext('2d');

          this.lineChart = new Chart(ctx, {
            type: 'scatter',
            data: {
              labels: xLabels,
              datasets: datasets,
            },
            options: {
              scales: {
                x: {
                  beginAtZero: true,
                },
              },
            },
          });
        } else {
          this.openSnackBar(
            'Aqui va un alert, que x4 debe ser mayor a x3',
            'close'
          );
        }
      } else {
        this.openSnackBar(
          'Aqui va un alert que x2 debe ser mayor a x1',
          'close'
        );
      }
    }
  }

  openSnackBar(message: string, action: string) {
    this._snackBar.open(message, action);
  }

  onKeyDown(event: KeyboardEvent) {
    // Permitir solo números, punto, retroceso y teclas de flecha
    const allowedChars = /[0-9\.]|Backspace|Tab|ArrowLeft|ArrowRight|ArrowUp|ArrowDown/;
    const key = event.key;
    if (!allowedChars.test(key)) {
      event.preventDefault();
    }
  }
}
