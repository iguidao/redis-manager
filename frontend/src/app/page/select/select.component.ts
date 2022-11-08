import { Component, OnInit, Input, Output, EventEmitter, ViewChild } from '@angular/core';

@Component({
  selector: 'app-select',
  templateUrl: './select.component.html',
  styleUrls: ['./select.component.css']
})
export class SelectComponent implements OnInit {
  selected!: string;
  @Input()

  // @Output() onSelect = new EventEmitter<string>();
  @Output() onSelect: EventEmitter<string> = new EventEmitter();
  constructor() {}

  ngOnInit() {
  }

  select(name: string) {
    
    if (name === this.selected) {
      return
    };

    this.onSelect.emit(name);

  }
}
