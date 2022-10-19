import { Component, Input, OnInit } from '@angular/core';
import { Item } from '../../Interfaces/Interfaces';

@Component({
  selector: 'app-item',
  templateUrl: './item.component.html',
  styleUrls: ['./item.component.scss'],
})
export class ItemComponent implements OnInit {
  @Input() item: Item | undefined;

  constructor() {}

  ngOnInit(): void {}
}
