import { TestBed } from '@angular/core/testing';

import { GuidaoadminService } from './guidaoadmin.service';

describe('GuidaoadminService', () => {
  let service: GuidaoadminService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GuidaoadminService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
