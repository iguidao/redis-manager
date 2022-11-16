import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DcodisComponent } from './dcodis.component';

describe('DcodisComponent', () => {
  let component: DcodisComponent;
  let fixture: ComponentFixture<DcodisComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DcodisComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DcodisComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
