import { Injectable } from "@angular/core";
import { APIResponse } from "../../domain/base/api_response";
import { Garment, GarmentValues } from "../../domain/stock/garment";

@Injectable({ providedIn: "root" })
export class StockService {
	private readonly root = "http://localhost:8080/stock"
	private readonly timer: number = 3000;
	private readonly requestConfig = { signal: AbortSignal.timeout(this.timer) }

	async get(id: number): Promise<APIResponse<Garment>> {
		// Initialize an empty Response from our backend API.
		let serverResponse: Response;

		// Attempt to connect and bail early if it fails.
		try {
			serverResponse = await fetch(`${this.root}/${id}`, this.requestConfig);
		} catch (error) {
			return new APIResponse<Garment>(
				false,
				"There was an unexpected error in the StockService layer (call get).",
			);
		}

		// Await the JSON body and return the APIResponse.
		let body: { data: Garment | string } = await serverResponse.json();
		return new APIResponse<Garment>(
			serverResponse.ok,
			body.data,
		)
	}

	async getMany(): Promise<APIResponse<Garment[]>> {
		// Initialize an empty Response from our backend API.
		let serverResponse: Response;

		// Attempt to connect and bail early if it fails.
		try {
			serverResponse = await fetch(`${this.root}`, this.requestConfig)
		} catch (error) {
			return new APIResponse<Garment[]>(
				false,
				"There was an unexpected error in the StockService layer (call getMany).",
			);
		}

		// Await the JSON body and return the APIResponse.
		let body: { data: Garment[] | string } = await serverResponse.json();
		return new APIResponse<Garment[]>(
			serverResponse.ok,
			body.data,
		)
	}

	async create(gv: GarmentValues): Promise<APIResponse<Garment>> {
		// Initialize an empty Response.
		let serverResponse: Response;

		// Attempt to connect and bail early if it fails.
		try {
			serverResponse = await fetch(`${this.root}`, { method: "POST", body: JSON.stringify(gv), ...this.requestConfig })
		} catch (error) {
			return new APIResponse<Garment>(
				false,
				"There was an unexpected error in the StockService layer (call get).",
			);
		}

		let body: { data: Garment | string } = await serverResponse.json();
		return new APIResponse<Garment>(
			serverResponse.ok,
			body.data,
		)
	}

	async update(id: number, gv: GarmentValues): Promise<APIResponse<Garment>> {
		// Initialize an empty Response.
		let serverResponse: Response;

		// Attempt to connect and bail early if it fails.
		try {
			serverResponse = await fetch(`${this.root}/${id}`, { method: "POST", body: JSON.stringify(gv), ...this.requestConfig })
		} catch (error) {
			return new APIResponse<Garment>(
				false,
				"There was an unexpected error in the StockService layer (call get).",
			);
		}

		let body: { data: Garment | string } = await serverResponse.json();
		return new APIResponse<Garment>(
			serverResponse.ok,
			body.data,
		)
	}

	async delete(id: number): Promise<APIResponse<Garment>> {
		// Initialize an empty Response.
		let serverResponse: Response;

		// Attempt to connect and bail early if it fails.
		try {
			serverResponse = await fetch(`${this.root}/${id}`, { method: "DELETE", ...this.requestConfig })
		} catch (error) {
			return new APIResponse<Garment>(
				false,
				"There was an unexpected error in the StockService layer (call get).",
			);
		}

		let body: { data: Garment | string } = await serverResponse.json();
		return new APIResponse<Garment>(
			serverResponse.ok,
			body.data,
		)
	}

}