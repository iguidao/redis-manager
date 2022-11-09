import { TestBed } from '@angular/core/testing';

import { RedisService } from './redis.service';

describe('RedisService', () => {
  let service: RedisService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RedisService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
