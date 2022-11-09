import { TestBed } from '@angular/core/testing';

import { CodisService } from './codis.service';

describe('CodisService', () => {
  let service: CodisService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CodisService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
