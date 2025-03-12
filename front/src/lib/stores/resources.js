import {writable} from "svelte/store";
import axios from "axios";

export const resources = writable([])
export const resourcesError = writable("")

const apiBaseUrl = "/config_api/resources"

/*
***** DATA FORMAT EXAMPLE
* {
*  "id": "6ce750e5-05b8-4c4c-8fee-d9e381dbf364",
*  "ucaId": 56529,
*  "name": "M1 - Tutorat L2"
* }
*/

export function getResources() {
    resourcesError.set("")
    axios.get(`${apiBaseUrl}`)
        .then((res) => {
            resources.set(res.data)
        })
        .catch((err) => {
            console.log("An error has occurred while retrieving resources")
            console.log(err)
            if(err.response?.data?.message){
                resourcesError.set(JSON.stringify(err.response.data.message))
            } else {
                resourcesError.set(JSON.stringify(err))
            }
        })
}

export function getResource(id) {
    resourcesError.set("")
    return axios.get(`${apiBaseUrl}/${id}`)
        .then((res) => {
            return Promise.resolve(res.data)
        })
        .catch((err) => {
            console.log("An error has occurred while retrieving resource")
            console.log(err)
            if(err.response?.data?.message){
                resourcesError.set(JSON.stringify(err.response.data.message))
            } else {
                resourcesError.set(JSON.stringify(err))
            }
            return Promise.reject(err)
        })
}

export function postResources(ucaId, name) {
    resourcesError.set("")
    let parsed = parseInt(ucaId)
    return axios.post(`${apiBaseUrl}`,{
        ucaId: isNaN(parsed) ? ucaId : parsed,
        name: name
    })
        .then((res) => {
            getResources()
            return Promise.resolve()
        })
        .catch((err) => {
            console.log("An error as occurred while posting resource")
            console.log(err)
            if(err.response?.data?.message){
                resourcesError.set(JSON.stringify(err.response.data.message))
            } else {
                resourcesError.set(JSON.stringify(err))
            }
            return Promise.reject(err)
        })
}

export function putResource(id, ucaId, name) {
    resourcesError.set("")
    let parsed = parseInt(ucaId)
    return axios.put(`${apiBaseUrl}/${id}`,{
        ucaId: isNaN(parsed) ? ucaId : parsed,
        name: name
    })
        .then((res) => {
            getResources()
            return Promise.resolve()
        })
        .catch((err) => {
            console.log("An error as occurred while putting resource")
            console.log(err)
            if(err.response?.data?.message){
                resourcesError.set(JSON.stringify(err.response.data.message))
            } else {
                resourcesError.set(JSON.stringify(err))
            }
            return Promise.reject(err)
        })
}

export function deleteResource(id) {
    resourcesError.set("")
    return axios.delete(`${apiBaseUrl}/${id}`)
        .then((res) => {
            getResources()
            return Promise.resolve()
        })
        .catch((err) => {
            console.log("An error as occurred while deleting resource")
            console.log(err)
            if(err.response?.data?.message){
                resourcesError.set(JSON.stringify(err.response.data.message))
            } else {
                resourcesError.set(JSON.stringify(err))
            }
            return Promise.reject(err)
        })
}