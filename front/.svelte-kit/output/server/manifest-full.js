export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set([".DS_Store","favicon.png","fontawesome/.DS_Store","fontawesome/css/brands.css","fontawesome/css/fontawesome.css","fontawesome/css/regular.css","fontawesome/css/solid.css","fontawesome/webfonts/fa-brands-400.ttf","fontawesome/webfonts/fa-brands-400.woff2","fontawesome/webfonts/fa-regular-400.ttf","fontawesome/webfonts/fa-regular-400.woff2","fontawesome/webfonts/fa-solid-900.ttf","fontawesome/webfonts/fa-solid-900.woff2","fontawesome/webfonts/fa-v4compatibility.ttf","fontawesome/webfonts/fa-v4compatibility.woff2","js/bootstrap.min.js","style/bootstrap.css","style/bootstrap_pulse.css","style/custom.css"]),
	mimeTypes: {".png":"image/png",".css":"text/css",".ttf":"font/ttf",".woff2":"font/woff2",".js":"text/javascript"},
	_: {
		client: {start:"_app/immutable/entry/start.CIwdh4z-.js",app:"_app/immutable/entry/app.NtNK2BXx.js",imports:["_app/immutable/entry/start.CIwdh4z-.js","_app/immutable/chunks/DkYErxwx.js","_app/immutable/chunks/BSISFcvB.js","_app/immutable/chunks/BYBiPWhc.js","_app/immutable/chunks/BUpOKEDO.js","_app/immutable/entry/app.NtNK2BXx.js","_app/immutable/chunks/BSISFcvB.js","_app/immutable/chunks/CApVEXk3.js","_app/immutable/chunks/Dy6sgHUi.js","_app/immutable/chunks/_Nf9tCAq.js","_app/immutable/chunks/BUpOKEDO.js","_app/immutable/chunks/BdwiA3Aa.js"],stylesheets:[],fonts:[],uses_env_dynamic_public:false},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js')),
			__memo(() => import('./nodes/2.js')),
			__memo(() => import('./nodes/3.js')),
			__memo(() => import('./nodes/4.js')),
			__memo(() => import('./nodes/5.js'))
		],
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 2 },
				endpoint: null
			},
			{
				id: "/alerts",
				pattern: /^\/alerts\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 3 },
				endpoint: null
			},
			{
				id: "/events",
				pattern: /^\/events\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 4 },
				endpoint: null
			},
			{
				id: "/resources",
				pattern: /^\/resources\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 5 },
				endpoint: null
			}
		],
		prerendered_routes: new Set([]),
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();
