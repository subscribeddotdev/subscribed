// import { paths } from "@@/constants";
// import { NextRequest, NextResponse } from "next/server";

// enum Constants {
//   CookieToken = "sbs_token",
// }

// export function authMiddleware(protectedRoutes: string[]) {
//   return (request: NextRequest) => {
//     const url = new URL(request.url);
//     const token = request.cookies.get(Constants.CookieToken);
//     const isProtected = protectedRoutes.find((route) => url.pathname.startsWith(route));

//     if (isProtected && !token) {
//       return NextResponse.redirect(new URL(paths.signin, request.url));
//     }

//     return NextResponse.next();
//   };
// }
