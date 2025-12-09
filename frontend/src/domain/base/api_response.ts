export class APIResponse<T> {
	public ok: boolean;
	public data: T | string;
	constructor(ok: boolean, data: T | string) {
		this.ok = ok;
		this.data = data;
	}
	setData(data: T | string) {
		this.data = data
	}
}