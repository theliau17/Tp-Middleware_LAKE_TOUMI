export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["favicon.png","fontawesome/css/brands.css","fontawesome/css/fontawesome.css","fontawesome/css/regular.css","fontawesome/css/solid.css","fontawesome/webfonts/fa-brands-400.ttf","fontawesome/webfonts/fa-brands-400.woff2","fontawesome/webfonts/fa-regular-400.ttf","fontawesome/webfonts/fa-regular-400.woff2","fontawesome/webfonts/fa-solid-900.ttf","fontawesome/webfonts/fa-solid-900.woff2","fontawesome/webfonts/fa-v4compatibility.ttf","fontawesome/webfonts/fa-v4compatibility.woff2","js/bootstrap.min.js","style/bootstrap.css","style/bootstrap_pulse.css","style/custom.css"]),
	mimeTypes: {".png":"image/png",".css":"text/css",".ttf":"font/ttf",".woff2":"font/woff2",".js":"text/javascript"},
	_: {
		client: {start:"_app/immutable/entry/start.C5VkjscK.js",app:"_app/immutable/entry/app.CtWl09_3.js",imports:["_app/immutable/entry/start.C5VkjscK.js","_app/immutable/chunks/BKgRMx3J.js","_app/immutable/chunks/BSISFcvB.js","_app/immutable/chunks/BYBiPWhc.js","_app/immutable/chunks/BUpOKEDO.js","_app/immutable/entry/app.CtWl09_3.js","_app/immutable/chunks/BSISFcvB.js","_app/immutable/chunks/CApVEXk3.js","_app/immutable/chunks/Dy6sgHUi.js","_app/immutable/chunks/_Nf9tCAq.js","_app/immutable/chunks/BUpOKEDO.js","_app/immutable/chunks/BdwiA3Aa.js"],stylesheets:[],fonts:[],uses_env_dynamic_public:false},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js'))
		],
		routes: [
			
		],
		prerendered_routes: new Set(["/","/alerts","/events","/resources"]),
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();
