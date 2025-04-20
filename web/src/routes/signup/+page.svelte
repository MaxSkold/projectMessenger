<script lang="ts">
    import RequiredLabel from "$lib/RequiredLabel.svelte";

    let email: string = '';
    let phone: string = '';
    let password: string = '';
    let confirmPassword: string = '';
    let message: string = '';
    let loading: boolean = false;


    function normalizePhone(phone: string): string {
        return phone.replace(/[^+\d]/g, '');
    }

    function isPasswordStrong(pw: string): boolean {
        return pw.length >= 8 &&
            /[A-Z]/.test(pw) &&
            /[a-z]/.test(pw) &&
            /\d/.test(pw);
        // /[!@#$%^&*(),.?":{}|<>]/.test(pw);
    }

    async function register() {
        message = '';

        if (password !== confirmPassword) {
            message = 'Passwords do not match';
            return;
        }

        if (!isPasswordStrong(password)) {
            message = 'Password must be at least 8 characters long and include uppercase, lowercase, number character';
            return;
        }

        loading = true;

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
                email = phone = password = confirmPassword = '';
            } else {
                const userExists = await res.json();

                if (userExists.error === 'user already exists') {
                    message = 'User with this email already exists';
                }
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
            <RequiredLabel forId="email" label="Email" required/>
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
            <RequiredLabel forId="phone" label="Phone"/>
            <input
                    id="phone"
                    type="tel"
                    bind:value={phone}
                    placeholder="+1 (123) 456-7890"
                    class="mt-1 w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    autocomplete="tel"
            />
        </div>

        <div>
            <RequiredLabel forId="password" label="Password" required/>
            <input
                    id="password"
                    type="password"
                    bind:value={password}
                    required
                    class="mt-1 w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    autocomplete="new-password"
            />
        </div>

        <div>
            <RequiredLabel forId="confirmPassword" label="Confirm Password" required/>
            <input
                    id="confirmPassword"
                    type="password"
                    bind:value={confirmPassword}
                    required
                    class="mt-1 w-full px-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    on:copy|preventDefault
                    on:cut|preventDefault
                    on:paste|preventDefault
                    on:contextmenu|preventDefault
                    on:selectstart|preventDefault
                    autocomplete="off"
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
