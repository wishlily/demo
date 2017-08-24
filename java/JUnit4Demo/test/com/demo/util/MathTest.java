package com.demo.util;

import static org.junit.Assert.*;

import org.junit.Test;

public class MathTest extends Math {

    @Test
    public void testFactorial() throws Exception {
        assertEquals(120, new Math().factorial(5));
    }

    @Test
    public void testAdd() {
        assertEquals(3, new Math().add(1, 2));
        assertEquals(0, new Math().add(1, -1));
        assertEquals(-3, new Math().add(-5, 2));
    }
}
