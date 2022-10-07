import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ItemsoverviewComponent } from './itemsoverview.component';

describe('ItemsoverviewComponent', () => {
  let component: ItemsoverviewComponent;
  let fixture: ComponentFixture<ItemsoverviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ItemsoverviewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ItemsoverviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
