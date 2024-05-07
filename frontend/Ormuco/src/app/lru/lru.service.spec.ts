import { TestBed } from '@angular/core/testing';

import { LruService } from './lru.service';

describe('LruService', () => {
  let service: LruService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(LruService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
