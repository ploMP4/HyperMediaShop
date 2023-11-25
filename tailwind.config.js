/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.templ", "./templates/**/*.templ"],
  theme: {
    extend: {
      gridTemplateRows: {
        "[auto,auto,ifr]": "auto auto ifr",
      },
    },
  },
  plugins: [require("@tailwindcss/aspect-ratio")],
};
