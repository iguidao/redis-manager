import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DredisComponent } from './dredis.component';

describe('DredisComponent', () => {
  let component: DredisComponent;
  let fixture: ComponentFixture<DredisComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DredisComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DredisComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
