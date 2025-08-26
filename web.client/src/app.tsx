import { useState } from "react";

const API_URL = "http://localhost:8080";

export function App() {
  const [password, setPassword] = useState("");

  const generatePassword = async () => {
    const response = await fetch(
      `${API_URL}/generate?sets=lowercase,uppercase,digits,special`
    );
    const data = await response.text();
    setPassword(data);
  };

  return (
    <>
      <button onClick={generatePassword}>Random password</button>
      {password && <p>Your new password is: {password}</p>}
    </>
  );
}
