"use client"; // Add this to the top of your file

import { useState } from "react";
import LoginForm from "@/app/login/LoginForm";
import SignupForm from "@/app/login/SignupForm";
import { Button } from "@/components/ui/button";
import Image from "next/image";

export default function AuthPage() {
  const [isSignUp, setIsSignUp] = useState(false);

  const handleToggle = () => {
    setIsSignUp((prev) => !prev);
  };

  return (
    <div className="min-h-screen grid grid-cols-1 lg:grid-cols-2">
      {/* Left Side - Image with Text Overlay */}
      <div className="relative hidden lg:block">
        <Image
          src="/images/login-bg.png"
          alt="Backend image"
          layout="fill"
          objectFit="cover"
          className="absolute inset-0"
        />
        <div className="absolute inset-0 bg-black opacity-50" />
        <div className="absolute inset-0 flex flex-col items-center justify-center text-center text-white p-8">
          <h2 className="text-4xl font-bold">Post once at everywhere</h2>
          <p className="mt-4 text-lg">
            Manage your tasks efficiently with PostIt’s.Without posting same post on every social media.
          </p>
        </div>
      </div>

      {/* Right Side - Login/SignUp Form */}
      <div className="flex flex-col items-center justify-center p-8 lg:p-16 bg-background text-foreground">
        <div className="max-w-md w-full">
          <h1 className="text-3xl font-bold mb-4 text-center">
            {isSignUp ? "Sign Up" : "Welcome to PostIt"}
          </h1>
          {isSignUp ? (
            <SignupForm onSignupSuccess={() => setIsSignUp(false)} />
          ) : (
            <LoginForm />
          )}

          <div className="relative my-6">
            <div className="absolute inset-0 flex items-center">
              <div className="w-full border-t border-gray-600"></div>
            </div>
            <div className="relative flex justify-center text-sm">
              <span className="bg-background px-2 text-gray-200">or</span>
            </div>
          </div>

          {/* Toggle between Login and Signup */}
          <Button
            variant="outline"
            className={`w-full ${isSignUp ? "bg-blue-600 text-white hover:bg-blue-700" : "bg-green-600 text-white hover:bg-green-700"}`}
            onClick={handleToggle}
          >
            {isSignUp ? "Back to Login" : "Sign Up"}
          </Button>
        </div>
      </div>
    </div>
  );
}