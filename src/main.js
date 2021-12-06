import RegisterComponents from './SvelteComponents.js';

import App from './App.svelte';
import Counter from './Counter.svelte';

let Components = {
	"app": App,
	"counter": Counter,
};

RegisterComponents(Components);
