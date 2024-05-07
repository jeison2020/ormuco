import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})

export class LibraryService {

  constructor(private http: HttpClient) { }

  getCompareVersions(version1: string, version2: string): Observable<any> {
    const url = `${environment.apiUrl}/compare/${version1}/${version2}`;
    return this.http.get(url);
  }
}
