import { TestBed } from '@angular/core/testing';

import { KodoService } from './kodo.service';

describe('KodoService', () => {
  let service: KodoService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(KodoService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
