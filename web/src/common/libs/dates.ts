import dayjs from "dayjs";
import localizedFormat from "dayjs/plugin/localizedFormat";
import relativeTime from "dayjs/plugin/relativeTime";

// Configure plugins
dayjs.extend(localizedFormat);
dayjs.extend(relativeTime);

export const dates = dayjs;
