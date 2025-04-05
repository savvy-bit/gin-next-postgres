"use client"
import { useRouter } from "next/navigation";
import { useState } from "react";
import { FaEnvelope, FaKey } from "react-icons/fa";
import { AuthCredentials } from "@/types/auth.type";
import InputText from "@/components/common/InputBox";
import Section from "@/components/common/Section";
import { Button } from "@/components/common/Button";
import { ROUTER } from "@/constants/common";


const LoginPage: React.FC = () => {
  const [credentials, setCredentials] = useState<AuthCredentials>({ email: "", password: "" })

  const router = useRouter();

  return (
    <Section className="w-[400px] m-auto">
      <InputText
        label="Email"
        labelIcon={<FaEnvelope size={16} />}
        value={credentials.email}
        onChange={(e) => setCredentials({ ...credentials, email: e.target.value })}
      />
      <InputText
        label="Password"
        labelIcon={<FaKey size={16} />}
        onChange={(e) => setCredentials({ ...credentials, password: e.target.value })}
        type="password"
      />
      <div className="flex justify-evenly gap-4 mt-4">
        <Button
          customClass="w-30"
          onClick={() => { }}
        >
          Reigster
        </Button>
        <Button
          customClass="text-nowrap"
          onClick={() => router.push(ROUTER.Login)}
          isPrimary={false}
        >
          Already have account
        </Button>
      </div>
    </Section >
  )
}

export default LoginPage;