import { startTransition, useActionState, useState } from "react";
import { AVAILABLE_SETS, generatePassword, type PasswordSet } from "./lib/api";

export function App() {
  const [selectedSets, setSelectedSets] = useState<PasswordSet[]>([]);
  const [password, handleGeneratePassword, isPending] = useActionState(
    async () => {
      const password = await generatePassword(selectedSets);
      return password;
    },
    null
  );

  const handleCheckboxChange = (set: PasswordSet) => {
    setSelectedSets((prev) =>
      prev.includes(set) ? prev.filter((s) => s !== set) : [...prev, set]
    );
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 px-4">
      <div className="bg-white rounded-xl shadow-lg p-8 w-full max-w-md">
        <h1 className="text-2xl font-bold mb-4 text-center">
          Password Generator
        </h1>
        <div className="flex flex-wrap gap-4 justify-center">
          {AVAILABLE_SETS.map((set) => (
            <label key={set} className="flex items-center gap-2">
              <input
                type="checkbox"
                checked={selectedSets.includes(set)}
                onChange={() => handleCheckboxChange(set)}
                className="accent-blue-600"
              />
              <span className="capitalize">{set}</span>
            </label>
          ))}
        </div>

        <button
          disabled={isPending}
          onClick={() => startTransition(handleGeneratePassword)}
          className="mt-4 w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded mb-6 transition-colors disabled:bg-blue-300 disabled:cursor-not-allowed"
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
