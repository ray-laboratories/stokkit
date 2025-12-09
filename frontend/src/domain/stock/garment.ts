import { Entity } from "../base/entity";

export class GarmentValues {
	public year: number;
	public description: string;
	public style: string;

	constructor(year: number, description: string, style: string) {
		this.year = year;
		this.description = description;
		this.style = style;
	}
}

// A "Garment" is both a set of GarmentValues
// and persistence details (or an Entity).
export type Garment = GarmentValues & Entity