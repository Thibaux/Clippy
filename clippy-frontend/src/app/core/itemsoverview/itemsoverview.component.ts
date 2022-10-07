import { Component, OnInit } from '@angular/core';
import { Item } from '../../Interfaces/Interfaces';

@Component({
  selector: 'app-itemsoverview',
  templateUrl: './itemsoverview.component.html',
  styleUrls: ['./itemsOverview.component.scss'],
})
export class ItemsoverviewComponent implements OnInit {
  items!: Item[];

  constructor() {}

  ngOnInit(): void {
    this.items = [
      {
        id: '123',
        name: 'name',
      },
    ];
    console.log(this.items);
  }
}
