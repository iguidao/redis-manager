import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OpredisComponent } from './opredis.component';

describe('OpredisComponent', () => {
  let component: OpredisComponent;
  let fixture: ComponentFixture<OpredisComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ OpredisComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OpredisComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
