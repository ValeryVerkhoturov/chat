/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{html,js}", "./images/**/*.svg"],
  theme: {
    extend: {
      fontSize: {
          "base": "14px",
      },
      keyframes: {
        fadeInUp: {
          '0%': { opacity: '0', transform: 'translateY(10px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
      },
      animation: {
        fadeInUp: 'fadeInUp 0.5s ease-out',
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}

