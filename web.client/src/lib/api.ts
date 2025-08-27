import invariant from "tiny-invariant";

const API_URL = import.meta.env.VITE_API_URL;
invariant(API_URL, "VITE_API_URL is not defined");

function makeRequest(path: string, options: RequestInit = {}) {
  return fetch(`${API_URL}${path}`, options);
}

export const AVAILABLE_SETS = ["lowercase", "uppercase", "digits", "special"];
export type PasswordSet = (typeof AVAILABLE_SETS)[number];

export async function generatePassword(
  sets: PasswordSet[] = ["lowercase", "uppercase", "digits", "special"]
) {
  const query = new URLSearchParams();

  if (sets.length > 0) {
    query.append("sets", sets.join(","));
  }

  const response = await makeRequest(
    `service.generator/generate?${query.toString()}`
  );

  if (!response.ok) {
    throw new Error("Failed to generate password");
  }

  const data = await response.text();
  return data;
}
