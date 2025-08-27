import invariant from "tiny-invariant";

const API_URL = import.meta.env.VITE_API_URL;
invariant(API_URL, "VITE_API_URL is not defined");

function makeRequest(path: string, options: RequestInit = {}) {
  return fetch(`${API_URL}${path}`, options);
}

export async function generatePassword() {
  const response = await makeRequest(
    "service.generator/generate?sets=lowercase,uppercase,digits,special"
  );

  if (!response.ok) {
    throw new Error("Failed to generate password");
  }

  const data = await response.text();
  return data;
}
