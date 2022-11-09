import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OpcodisComponent } from './opcodis.component';

describe('OpcodisComponent', () => {
  let component: OpcodisComponent;
  let fixture: ComponentFixture<OpcodisComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ OpcodisComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OpcodisComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
