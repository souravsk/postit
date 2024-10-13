"use client"; // Add this to the top of your file

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Toast, ToastProvider, ToastViewport } from "@/components/ui/toast"; // Adjust import based on your ShadCN setup

// Define the props interface
interface SignupFormProps {
    onSignupSuccess: () => void;
  }

export default function SignupForm({ onSignupSuccess } : SignupFormProps) {
  const [user, setUser] = useState("");
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [successMessage, setSuccessMessage] = useState("");

  const handleSignup = async (event: React.FormEvent) => {
    event.preventDefault();

    if (password !== confirmPassword) {
      alert("Passwords do not match!");
      return;
    }

    const response = await fetch("http://localhost:8080/user/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name, email, password, user }),
    });

    const data = await response.json();
    if (response.ok) {
      setSuccessMessage("Created user successfully");
      setTimeout(() => {
        setSuccessMessage(""); // Clear the message after a short delay
        onSignupSuccess(); // Call the callback to switch to login
      }, 2000); // Adjust the timeout as needed
    } else {
      console.error("Signup failed:", data);
    }
  };

  return (
    <ToastProvider>
      <form className="space-y-6" onSubmit={handleSignup}>
        <div>
          <label htmlFor="user" className="block text-sm font-medium text-gray-200">
            User
          </label>
          <Input
            type="text"
            id="user"
            value={user}
            onChange={(e) => setUser(e.target.value)}
            placeholder="souravk"
            className="mt-1 block w-full"
            required
          />
        </div>
        <div>
          <label htmlFor="name" className="block text-sm font-medium text-gray-200">
            Name
          </label>
          <Input
            type="text"
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Your Name"
            className="mt-1 block w-full"
            required
          />
        </div>
        <div>
          <label htmlFor="email" className="block text-sm font-medium text-gray-200">
            Email
          </label>
          <Input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="you@example.com"
            className="mt-1 block w-full"
            required
          />
        </div>
        <div>
          <label htmlFor="password" className="block text-sm font-medium text-gray-200">
            Password
          </label>
          <Input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="********"
            className="mt-1 block w-full"
            required
          />
        </div>
        <div>
          <label htmlFor="confirmPassword" className="block text-sm font-medium text-gray-200">
            Confirm Password
          </label>
          <Input
            type="password"
            id="confirmPassword"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            placeholder="********"
            className="mt-1 block w-full"
            required
          />
        </div>
        <Button className="w-full" type="submit">
          Submit
        </Button>
      </form>

      {successMessage && (
        <Toast>
          {successMessage}
        </Toast>
      )}

      <ToastViewport />
    </ToastProvider>
  );
}