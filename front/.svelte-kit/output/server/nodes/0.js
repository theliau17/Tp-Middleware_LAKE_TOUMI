import * as server from '../entries/pages/_layout.server.js';

export const index = 0;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_layout.svelte.js')).default;
export { server };
export const server_id = "src/routes/+layout.server.js";
export const imports = ["_app/immutable/nodes/0.BrZRVdXK.js","_app/immutable/chunks/Dy6sgHUi.js","_app/immutable/chunks/BSISFcvB.js","_app/immutable/chunks/CHrprlqx.js","_app/immutable/chunks/BKgRMx3J.js","_app/immutable/chunks/BYBiPWhc.js","_app/immutable/chunks/BUpOKEDO.js"];
export const stylesheets = ["_app/immutable/assets/0.C1IPNx-7.css"];
export const fonts = [];
