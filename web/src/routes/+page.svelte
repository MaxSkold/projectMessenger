<script lang="ts">
    let email = '';
    let phone = '';
    let password = '';
    let message = '';
    let loading = false;

    async function register() {
        message = '';
        loading = true;

        const normalizePhone = (phone: string) =>
            phone.replace(/[^+\d]/g, '');

        const normalizedPhone = normalizePhone(phone);

        try {
            const res = await fetch('http://localhost:8080/api/signup', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    email,
                    phone_number: normalizedPhone,
                    password
                })
            });

            if (res.ok) {
                message = `User is registered successfully`;
                email = phone = password = '';
            } else {
                const err = await res.json();
                message = `Error: ${err.error || res.statusText}`;
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
            on:submit|preventDefault={register}
            class="w-full max-w-md bg-white p-8 rounded-2xl shadow-xl space-y-6"
            aria-labelledby="signup-title"
    >
        <h2 id="signup-title" class="text-2xl font-bold text-center text-gray-800">Sign Up</h2>

        <div>
            <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
            <input
                    id="email"
                    type="email"
                    bind:value={email}
                    required
                    placeholder="user@example.com"
                    class="mt-1 w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    autocomplete="email"
            />
        </div>

        <div>
            <label for="phone" class="block text-sm font-medium text-gray-700">Phone Number</label>
            <input
                    id="phone"
                    type="tel"
                    bind:value={phone}
                    required
                    placeholder="+1 (123) 456-7890"
                    class="mt-1 w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    autocomplete="tel"
            />
        </div>

        <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
            <input
                    id="password"
                    type="password"
                    bind:value={password}
                    required
                    minlength="8"
                    class="mt-1 w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    autocomplete="new-password"
            />
        </div>

        <button
                type="submit"
                class="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded-lg transition duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
                disabled={loading}
                aria-busy={loading}
        >
            {loading ? 'Loading...' : 'Register'}
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
