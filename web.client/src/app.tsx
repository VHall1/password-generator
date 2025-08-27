import { useState } from "react";
import { generatePassword } from "./lib/api";

export function App() {
  const [password, setPassword] = useState("");

  const handleGeneratePassword = async () => {
    const password = await generatePassword();
    setPassword(password);
  };

  return (
    <>
      <button onClick={handleGeneratePassword}>Random password</button>
      {password && <p>Your new password is: {password}</p>}
    </>
  );
}
