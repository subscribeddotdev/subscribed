import "@@/common/styles/globals.css";
import { Theme } from "@radix-ui/themes";
import "@radix-ui/themes/styles.css";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { HelmetProvider } from "react-helmet-async";
import {
  createBrowserRouter,
  Navigate,
  RouterProvider,
} from "react-router-dom";
import DashboardHomePage from "./pages/Dashboard/DashboardHomePage.tsx";
import EventTypesPage from "./pages/Dashboard/EventTypesPage.tsx";
import SignInPage from "./pages/signin.tsx";
import SignUpPage from "./pages/signup.tsx";

const router = createBrowserRouter([
  { path: "/", element: <Navigate to="/signin" /> },
  { path: "/signin", element: <SignInPage /> },
  { path: "/signup", element: <SignUpPage /> },
  {
    path: "/:environment",
    element: <DashboardHomePage />,
  },
  {
    path: "/:environment/event-types",
    element: <EventTypesPage />,
  },
]);

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <HelmetProvider>
      <Theme accentColor="indigo" appearance="dark">
        <RouterProvider router={router} />
      </Theme>
    </HelmetProvider>
  </StrictMode>
);
