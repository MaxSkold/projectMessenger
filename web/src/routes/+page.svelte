<script lang="ts">
    let phoneOrEmail = '';

    let password = '';
    let message = '';
    let loading = false;

    function isEmail(input: string): boolean {
        return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(input);
    }

    function sanitizePhone(phone: string): string {
        return phone.replace(/[^+\d]/g, '');
    }

    async function login() {
        message = '';
        loading = true;

        const payload: Record<string, string> = {
            password
        };

        if (isEmail(phoneOrEmail)) {
            payload.email = phoneOrEmail;
        } else {
            payload.phone_number = sanitizePhone(phoneOrEmail);
        }

        try {
            const res = await fetch('http://localhost:8080/api/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload)
            });

            if (res.ok) {
                message = '✅ Logged in successfully';
                // редирект в чат или сохраняем токен
            } else {
                const err = await res.json();
                message = `${err.error || res.statusText}`;
            }
        } catch (err) {
            message = 'Connection server error';
            console.error(err);
        } finally {
            loading = false;
        }
    }
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-100 p-4">
    <form
            on:submit|preventDefault={login}
            class="w-full max-w-md bg-white p-8 rounded-2xl shadow-xl space-y-6"
            aria-labelledby="login-title"
    >
        <h2 id="hello-title" class="text-2xl font-bold text-center text-blue-800">Welcome Back!</h2>
        <h2 id="login-title" class="text-3xl font-bold text-center text-gray-800">Log In</h2>
        <div>
            <label for="identifier" class="block text-sm font-medium text-gray-700">Phone or Email</label>
            <input
                    id="identifier"
                    type="text"
                    bind:value={phoneOrEmail}
                    required
                    placeholder="+1 (123) 456-7890 or user@example.com"
                    class="mt-1 w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    autocomplete="username"
            />
        </div>

        <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
            <input
                    id="password"
                    type="password"
                    bind:value={password}
                    required
                    class="mt-1 w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    autocomplete="current-password"
            />
        </div>

        <div class="flex justify-between text-sm text-blue-600">
            <a href="/forgot" class="hover:underline">Forgot password?</a>
            <a href="/signup" class="hover:underline">New user?</a>
        </div>

        <button
                type="submit"
                class="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded-lg transition duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
                disabled={loading}
                aria-busy={loading}
        >
            {loading ? 'Logging in...' : 'Log In'}
        </button>

        {#if message}
            <p
                    role="alert"
                    class="text-center text-sm mt-2 {message.startsWith('✅') ? 'text-green-600' : 'text-red-600'}"
            >
                {message}
            </p>
        {/if}
    </form>
</div>
