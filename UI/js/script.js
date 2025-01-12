// origin is same for api and UI
const apiBase = "/api/v1";

// Helper function to show messages
const showMessage = (message, type = "error") => {
  const msgEl = document.getElementById("message");
  msgEl.textContent = message;
  msgEl.classList.remove("hidden", "text-red-500", "text-green-500");
  msgEl.classList.add(type === "success" ? "text-green-500" : "text-red-500");
};

// Login
const loginForm = document.getElementById("loginForm");
if (loginForm) {
  loginForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    try {
      const res = await fetch(`${apiBase}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: `email=${email}&password=${password}`,
        credentials: "include",
      });
      const data = await res.json();
      if (res.ok) {
        location.href = "user.html";
      } else {
        showMessage(data.error);
      }
    } catch (error) {
      showMessage("An error occurred. Please try again.");
    }
  });
}

// Register
const registerForm = document.getElementById("registerForm");
if (registerForm) {
  registerForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const name = document.getElementById("name").value;
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    try {
      const res = await fetch(`${apiBase}/register`, {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: `name=${name}&email=${email}&password=${password}`,
      });
      const data = await res.json();
      if (res.ok) {
        showMessage("Registration successful!", "success");
        registerForm.reset();
      } else {
        showMessage(data.error);
      }
    } catch (error) {
      showMessage("An error occurred. Please try again.");
    }
  });
}

// User Page
const userPage = document.getElementById("userData");
if (userPage) {
  const logoutButton = document.getElementById("logout");
  logoutButton.addEventListener("click", async () => {
    await fetch(`${apiBase}/logout`, { credentials: "include" });
    location.href = "index.html";
  });

  (async () => {
    try {
      const res = await fetch(`${apiBase}/user`, { credentials: "include" });
      const data = await res.json();
      if (res.ok) {
        userPage.textContent = `Logged in as: ${data.name} (${data.email})`;
      } else {
        location.href = "index.html";
      }
    } catch {
      location.href = "index.html";
    }
  })();
}
