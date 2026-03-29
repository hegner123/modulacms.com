# The difference between slots, instance swaps, and variants

Before you Start

Who can use this feature

Available in Figma Design on all plans

Anyone with can edit access to a design file can add slots to components.

New to creating slots? Learn how slots brings flexibility to components.

### Slots vs instance swap

When it comes to building assets from a design system, there is a spectrum of needs in the amount of freedom or guardrails an asset needs.

On one end of the spectrum are components that require rigidity and more guardrails around how much you can configure the asset. Instance swapping is great for this need.

For example, let's say we have a browser tab component that is limited to one website icon at a time which must stay fixed to the right of the tab text. Also, designers should be able to swap the icon out. Instance swap allows you to limit a container to at most _one_ instance at a time and prevents designers from making layout changes to nested layers, such as dimensions and position.

On the other end of the spectrum we have components that require more flexibility and freeform designs while still maintaining adherence to design system guidelines. Slots is a great solution for this.

For example, let's say we have modals used across various ares in our product for messaging, settings, and confirmation. We want various properties—like padding, border radius, and drop shadows—to remain consistent across all permutations while giving designers the freedom to customize it with the type of content that they need for their particular use case and still receive updates if those properties are updated.

### Slots vs variants

Both slots and variants are types of component properties that broaden the number of variations and layouts that a reusable and customizable component can have.

While variants would previously be used by teams as workaround to replicate slots, especially for assets that required repeating items, the primary function of variants is to manage varying states or types of an asset.

For example, a button may have different action states—like default, hover, clicked, and disabled—so you'd create a variant for each state. In addition to action states, the button might also have different sizes based on the device type it's being used—like mobile and desktop.

Slots, on the other hand, are used to increase the flexibility of a base component by defining an area where you can add, edit, and arrange content as desired.

**Tip**: If you are currently using variants as a workaround to imitate slots, learn how to migrate to using to Figma's native slot feature.
