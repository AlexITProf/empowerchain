/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],

  theme: {
    extend: {
      colors: {
        greenPrimary: "#00C131",
        greenMedium: "#00B82E",
        greenDark: "#007A1F",
        greenTabs: "#00E33A",
        black: "rgba(0, 0, 0, 1)",
        lightBlack: "rgba(255, 255, 255, 0.06)",
        dropdownBlack: "rgba(33, 33, 33, 1)",
        white: "#FFFFFF",
        whiteLight: "#F6F5F5",
        darkGray: "rgba(255, 255, 255, 0.19)",
        lightGray: "rgba(255, 255, 255, 0.1)",
        buttonGray: "rgba(74, 74, 74, 1)",
        mediumGray: "rgba(255, 255, 255, 0.13)",
        textGray: "rgba(255, 255, 255, 0.6)",
        textLightGray: "rgba(155, 155, 155, 1)",
        textPlaceHolderGray: "rgba(230, 230, 230, 1)",
        dropDownText: "rgba(255, 255, 255, 0.6)",
        navBarBlack: "rgba(0, 0, 0, 0.95)",
        borderBlack: "rgba(255, 255, 255, 0.03)",
        dividerGray: "rgba(255, 255, 255, 0.08)",
        textWhite: "rgba(255, 255, 255, 0.8)",
        modalBackground: "rgba(0, 0, 0, 0.82)",
        modalBorder: "rgba(255, 255, 255, 0.07)",
        dashedBorderBlack: "rgba(255, 255, 255, 0.15)",
        borderGray: "rgba(255, 255, 255, 0.14)",
        buttonBlack: "rgba(255, 255, 255, 0.1)",
        paginationBlack: "rgb(155, 155, 155, 0.2)",
        avatarBlack: "#212121",
        avatarBorder: "rgba(255, 255, 255, 0.23)",
        darkGreen: "rgba(0, 227, 58, 0.13)",
        auctionBackground: "rgba(255, 255, 255, 0.14)",
        subTextGray: "rgba(255, 255, 255, 0.5)",
        dropdownGray: "#575757",
        dropdownGreen: "#1D5E0D",
        subLabel: "#9B9B9B",
        textInfoGray: "#B4B4B4",
        tabGray: "#3E3E3E",
        scrollButtonBackground: "rgba(255, 255, 255, 0.38)",
      },
      fontSize: {
        title32: "32px",
        title26: "26px",
        title16: "16px",
        title24: "24px",
        title40: "40px",
        title64: "64px",
        title12: "12px",
        title13: "13px",
        title11: "11px",
        title14: "14px",
        title15: "15px",
        title18: "18px",
        title38: "38px",
      },
      fontFamily: {
        openSans: ["Open Sans", "sans-serif"],
        Inter: ["Inter", "sans-serif"],
      },
      borderRadius: {
        sm: "15px",
        lg: "20px",
        xl: "50px",
      },
    },
  },
  daisyui: {
    styled: true,
    themes: false,
    base: false,
    utils: true,
    logs: true,
    rtl: false,
    prefix: "",
    darkTheme: "dark",
  },
  plugins: [require("daisyui")],
};
