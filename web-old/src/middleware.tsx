import { authMiddleware } from "./modules/Auth/authMiddleware";

export default authMiddleware(["/dashboard"]);

export const config = {
  matcher: ["/((?!.*\\..*|_next).*)", "/", "/(api|trpc)(.*)"],
};
