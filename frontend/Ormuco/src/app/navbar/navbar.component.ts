import { Component } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [
    RouterLink,
    RouterLinkActive,
    CommonModule,
  ],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.scss',
  providers: [],
})
export class NavbarComponent {
  constructor(
  ) {}

  navbar: Array<NavBar> = [
    { id: 'overlap', name: 'Overlap', routerLink: '/overlap' },
    { id: 'library', name: 'Software Library', routerLink: '/library' },
    { id: 'lru', name: 'Geo Distributed', routerLink: '/lru' },
  ];
}

export interface NavBar {
  id: string;
  name: string;
  routerLink: string;
}
