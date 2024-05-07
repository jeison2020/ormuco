import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class LibraryService {

  constructor(private http: HttpClient) { }
  
  compareVersions(version1: string, version2: string): Observable<any> {
    return this.http.get(`/api/v1/compare/${version1}/${version2}`);
  }
}
