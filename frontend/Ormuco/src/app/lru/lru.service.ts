import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Observable, catchError, throwError } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class LruService {
  constructor(private http: HttpClient) {}

  sendData(data: object): Observable<unknown> {
    console.log("data", data);
    const url = `${environment.apiUrl}/LRU`;
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
         return this.http.post(url, data, { headers });
  }

  getAllData(): Observable<any> {
    const url = `${environment.apiUrl}/LRU`;
    return this.http.get(url);
  }
}