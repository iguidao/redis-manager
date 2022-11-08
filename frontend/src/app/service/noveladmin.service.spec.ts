import { TestBed } from '@angular/core/testing';

import { NoveladminService } from './noveladmin.service';

describe('NoveladminService', () => {
  let service: NoveladminService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(NoveladminService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
