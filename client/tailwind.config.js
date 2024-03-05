/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        "UI-white": "#F9F9FD",
        "UI-black": "#1E1E1E",
        "Unimet-primary": {
          50: "#fffaec",
          100: "#fff4d3",
          200: "#ffe6a5",
          300: "#ffd26d",
          400: "#ffb332",
          500: "#ff9a0a",
          600: "#ff8200",
          700: "#cc5f02",
          800: "#a1490b",
          900: "#823e0c",
          950: "#461d04",
        },
        "primary-purple": {
          gradient: "#2C73D2",
          50: "#f8f7fc",
          100: "#f2f0fa",
          200: "#e7e4f6",
          300: "#d1c9ed",
          400: "#b8a7e2",
          500: "#9c82d4",
          600: "#8660c3",
          700: "#7953b2",
          800: "#654497",
          900: "#553b7d",
          950: "#382654",
        },
        "primary-blue": {
          gradient: "#5CA1FE",
          50: "#f0f8ff",
          100: "#e0f1fe",
          200: "#b9e3fe",
          300: "#7ccefd",
          400: "#36b7fa",
          500: "#0c9eeb",
          600: "#0081cf",
          700: "#0163a3",
          800: "#065486",
          900: "#0b466f",
          950: "#072d4a",
        },
        "primary-green": {
          gradient: "#008E9B",
          50: "#eefffa",
          100: "#c5fff1",
          200: "#8bffe4",
          300: "#4afed6",
          400: "#15ecc2",
          500: "#00d0ab",
          600: "#00a88c",
          700: "#008f7a",
          800: "#06695c",
          900: "#0a574c",
          950: "#003530",
        },
      },
    },
  },
  plugins: [],
};