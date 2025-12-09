import { Component, inject, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { StockService } from '../services/stock/service';
import { APIResponse } from '../domain/base/api_response';
import { Garment } from '../domain/stock/garment';

@Component({
	selector: 'app-root',
	imports: [RouterOutlet],
	templateUrl: './app.html',
	styleUrl: './app.css'
})
export class App {
	protected readonly title = signal('frontend');
	private readonly stockService = inject(StockService);

	async ngOnInit() {
		let resp: APIResponse<Garment[]> = await this.stockService.getMany();
		console.warn(JSON.stringify(resp));
	}
}
