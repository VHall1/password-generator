import { startTransition, useActionState } from "react";
import { generatePassword } from "./lib/api";

export function App() {
  const [password, handleGeneratePassword, isPending] = useActionState(
    async () => {
      const password = await generatePassword();
      return password;
    },
    null
  );

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 px-4">
      <div className="bg-white rounded-xl shadow-lg p-8 w-full max-w-md">
        <h1 className="text-2xl font-bold mb-4 text-center">
          Password Generator
        </h1>
        <button
          disabled={isPending}
          onClick={() => startTransition(() => handleGeneratePassword())}
          className="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded mb-6 transition-colors disabled:bg-blue-300 disabled:cursor-not-allowed"
        >
          Generate Password
        </button>

        {password && (
          <div className="text-center">
            <p className="text-gray-700 mb-2">Your new password is:</p>
            <span className="font-mono text-lg bg-gray-100 px-3 py-2 rounded border border-gray-300 inline-block">
              {password}
            </span>
          </div>
        )}
      </div>
    </div>
  );
}
