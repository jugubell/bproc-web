/*
 * File: index.js
 * Project: bproc-web
 * Last modified: 2025-11-18 23:11
 *
 * This file: index.js is part of BProC-WEB project.
 *
 * BProC-WEB is free software: you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published
 * by the Free Software Foundation, either version 2 of the License,
 * or (at your option) any later version.
 *
 * BProC-WEB is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty
 * of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with BProC-WEB. If not, see <https://www.gnu.org/licenses/>.
 *
 * Copyright (C) 2025 Jugurtha Bellagh
 */

// place files you want to import through the `$lib` alias in this folder.

import axios from 'axios';
import { editorDataStore } from '$lib/state.svelte.js';

const baseUrl = import.meta.env.VITE_BASE_URL || ''
const apiPrefix = "api"
const host = `${baseUrl}/${apiPrefix}`

export async function apiGet(url) {
	try {
		let res = await axios.get(`${host}/${url}`);
		console.log(res);
		return res;
	} catch (error) {
		console.log(error);
		return error;
	}
}

export async function apiPost(url, data) {
	try {
		return await axios.post(`${host}/${url}`, data);
	} catch (error) {
		console.log(error);
		return error;
	}
}

export function getCompileType() {
	const typeGroup = document.getElementById("compileTypeRadio");
	for(const inputDiv of typeGroup.querySelectorAll("div")) {
		const input = inputDiv.querySelector("input");
		if(input && input.checked) {
			return input.id.toLowerCase().replace(/cmp/g, "");
		}
	}
	return null;
}

export function notImplemented() {
	alert("Feature not implemented yet!");
}