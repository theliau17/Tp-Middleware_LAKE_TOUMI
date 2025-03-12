

export const index = 2;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/2.D2DfpwKv.js","_app/immutable/chunks/Dy6sgHUi.js","_app/immutable/chunks/BSISFcvB.js","_app/immutable/chunks/CHrprlqx.js"];
export const stylesheets = [];
export const fonts = [];
