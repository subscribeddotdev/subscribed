import { SignInPayload } from "@@/common/libs/backendapi/client";
import { LAST_CHOSEN_ENVIRONMENT } from "@@/constants";
import Cookies from "js-cookie";
import { useEffect, useState } from "react";

interface SignedInUserDetails {
  id: string;
  firstName: string;
  lastName: string;
  email: string;
}

export function storeTokenOnTheClient(token: string) {
  localStorage.setItem("sbs_token", token);
  Cookies.set("sbs_token", token, { expires: 7 });
}

export function retrieveTokenFromTheClient() {
  return localStorage.getItem("sbs_token");
}

export function retrieveTokenFromCookies(token: string) {
  return Cookies.get("sbs_token");
}

export function clearTokenFromCurrentSession() {
  localStorage.removeItem("sbs_token");
  localStorage.removeItem("sbs_user_details");
  localStorage.removeItem(LAST_CHOSEN_ENVIRONMENT);
  Cookies.remove("sbs_token");
}

export function storeUserDetails(details: SignInPayload) {
  const data: SignedInUserDetails = {
    id: details.id,
    firstName: details.first_name,
    lastName: details.last_name,
    email: details.email,
  };

  localStorage.setItem("sbs_user_details", JSON.stringify(data));
}

export function retrieveUserDetails(): SignedInUserDetails | null {
  const data = localStorage.getItem("sbs_user_details");
  return data ? (JSON.parse(data) as SignedInUserDetails) : null;
}

export function useUserDetails() {
  const [details, setDetails] = useState<SignedInUserDetails | null>();

  useEffect(() => {
    setDetails(retrieveUserDetails());
  }, []);

  return { details };
}
