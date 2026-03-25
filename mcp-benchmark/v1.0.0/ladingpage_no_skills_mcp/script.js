/* ========================================
   ARCTIS DEPLOY — INTERACTIONS
   ======================================== */

(function () {
  'use strict';

  // --- Scroll-triggered animations (IntersectionObserver) ---
  function initScrollAnimations() {
    const elements = document.querySelectorAll('[data-animate]');
    if (!elements.length) return;

    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            const el = entry.target;
            const delay = parseInt(el.dataset.delay || '0', 10);
            setTimeout(() => {
              el.classList.add('animated');
            }, delay);
            observer.unobserve(el);
          }
        });
      },
      {
        threshold: 0.15,
        rootMargin: '0px 0px -40px 0px',
      }
    );

    elements.forEach((el) => observer.observe(el));
  }

  // --- Counter animation for stats ---
  function initCounterAnimation() {
    const counters = document.querySelectorAll('[data-count]');
    if (!counters.length) return;

    function animateCounter(el) {
      const target = parseFloat(el.dataset.count);
      const isDecimal = target % 1 !== 0;
      const duration = 1600;
      const startTime = performance.now();

      function update(currentTime) {
        const elapsed = currentTime - startTime;
        const progress = Math.min(elapsed / duration, 1);
        const eased = 1 - Math.pow(1 - progress, 3);
        const current = target * eased;

        if (isDecimal) {
          el.textContent = current.toFixed(2);
        } else {
          el.textContent = Math.floor(current).toLocaleString('pt-BR');
        }

        if (progress < 1) {
          requestAnimationFrame(update);
        } else {
          if (isDecimal) {
            el.textContent = target.toFixed(2);
          } else {
            el.textContent = target.toLocaleString('pt-BR');
          }
        }
      }

      requestAnimationFrame(update);
    }

    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            animateCounter(entry.target);
            observer.unobserve(entry.target);
          }
        });
      },
      { threshold: 0.5 }
    );

    counters.forEach((el) => observer.observe(el));
  }

  // --- Navigation scroll state ---
  function initNavScroll() {
    const nav = document.getElementById('nav');
    if (!nav) return;

    let ticking = false;

    function updateNav() {
      if (window.scrollY > 20) {
        nav.classList.add('nav--scrolled');
      } else {
        nav.classList.remove('nav--scrolled');
      }
      ticking = false;
    }

    window.addEventListener('scroll', () => {
      if (!ticking) {
        requestAnimationFrame(updateNav);
        ticking = true;
      }
    }, { passive: true });

    updateNav();
  }

  // --- Mobile menu toggle ---
  function initMobileMenu() {
    const toggle = document.getElementById('mobileToggle');
    const menu = document.getElementById('mobileMenu');
    if (!toggle || !menu) return;

    toggle.addEventListener('click', () => {
      toggle.classList.toggle('active');
      menu.classList.toggle('active');
    });

    // Close menu when clicking a link
    const links = menu.querySelectorAll('a');
    links.forEach((link) => {
      link.addEventListener('click', () => {
        toggle.classList.remove('active');
        menu.classList.remove('active');
      });
    });

    // Close menu on escape
    document.addEventListener('keydown', (e) => {
      if (e.key === 'Escape' && menu.classList.contains('active')) {
        toggle.classList.remove('active');
        menu.classList.remove('active');
      }
    });
  }

  // --- Smooth scroll for anchor links ---
  function initSmoothScroll() {
    document.querySelectorAll('a[href^="#"]').forEach((anchor) => {
      anchor.addEventListener('click', function (e) {
        const href = this.getAttribute('href');
        if (href === '#') return;

        const target = document.querySelector(href);
        if (!target) return;

        e.preventDefault();
        const navHeight = document.getElementById('nav')?.offsetHeight || 72;
        const targetPos = target.getBoundingClientRect().top + window.scrollY - navHeight - 20;

        window.scrollTo({
          top: targetPos,
          behavior: 'smooth',
        });
      });
    });
  }

  // --- Subtle parallax on hero glow elements ---
  function initHeroParallax() {
    const glows = document.querySelectorAll('.hero__glow');
    if (!glows.length) return;

    // Only on desktop
    if (window.matchMedia('(max-width: 768px)').matches) return;

    let ticking = false;

    window.addEventListener('mousemove', (e) => {
      if (ticking) return;
      ticking = true;

      requestAnimationFrame(() => {
        const x = (e.clientX / window.innerWidth - 0.5) * 2;
        const y = (e.clientY / window.innerHeight - 0.5) * 2;

        glows.forEach((glow, i) => {
          const factor = (i + 1) * 8;
          glow.style.transform = `translate(${x * factor}px, ${y * factor}px)`;
        });

        ticking = false;
      });
    }, { passive: true });
  }

  // --- Init all ---
  document.addEventListener('DOMContentLoaded', () => {
    initScrollAnimations();
    initCounterAnimation();
    initNavScroll();
    initMobileMenu();
    initSmoothScroll();
    initHeroParallax();
  });
})();
