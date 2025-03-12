import { X as escape_html } from "./index.js";
import "clsx";
import { w as writable } from "./index2.js";
import axios from "axios";
const replacements = {
  translate: /* @__PURE__ */ new Map([
    [true, "yes"],
    [false, "no"]
  ])
};
function attr(name, value, is_boolean = false) {
  if (value == null || !value && is_boolean || value === "" && name === "class") return "";
  const normalized = name in replacements && replacements[name].get(value) || value;
  const assignment = is_boolean ? "" : `="${escape_html(normalized, true)}"`;
  return ` ${name}${assignment}`;
}
const SvelteMap = globalThis.Map;
const resources = writable([]);
const resourcesError = writable("");
const apiBaseUrl = "/config_api/resources";
function getResources() {
  resourcesError.set("");
  axios.get(`${apiBaseUrl}`).then((res) => {
    resources.set(res.data);
  }).catch((err) => {
    console.log("An error has occurred while retrieving resources");
    console.log(err);
    if (err.response?.data?.message) {
      resourcesError.set(JSON.stringify(err.response.data.message));
    } else {
      resourcesError.set(JSON.stringify(err));
    }
  });
}
function getResource(id) {
  resourcesError.set("");
  return axios.get(`${apiBaseUrl}/${id}`).then((res) => {
    return Promise.resolve(res.data);
  }).catch((err) => {
    console.log("An error has occurred while retrieving resource");
    console.log(err);
    if (err.response?.data?.message) {
      resourcesError.set(JSON.stringify(err.response.data.message));
    } else {
      resourcesError.set(JSON.stringify(err));
    }
    return Promise.reject(err);
  });
}
function deleteResource(id) {
  resourcesError.set("");
  return axios.delete(`${apiBaseUrl}/${id}`).then((res) => {
    getResources();
    return Promise.resolve();
  }).catch((err) => {
    console.log("An error as occurred while deleting resource");
    console.log(err);
    if (err.response?.data?.message) {
      resourcesError.set(JSON.stringify(err.response.data.message));
    } else {
      resourcesError.set(JSON.stringify(err));
    }
    return Promise.reject(err);
  });
}
export {
  SvelteMap as S,
  attr as a,
  resourcesError as b,
  deleteResource as d,
  getResource as g,
  resources as r
};
