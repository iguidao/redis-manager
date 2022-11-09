import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CodisComponent } from './codis.component';

describe('CodisComponent', () => {
  let component: CodisComponent;
  let fixture: ComponentFixture<CodisComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CodisComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CodisComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
