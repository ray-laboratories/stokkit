import Aurelia from 'aurelia';
import MainApp from "./main-app/main-app";
import { RouterConfiguration} from "@aurelia/router";

Aurelia
	.register(RouterConfiguration.customize({
		useUrlFragmentHash: false,
		historyStrategy: 'push',
	}))
	.app(MainApp)
	.start();
