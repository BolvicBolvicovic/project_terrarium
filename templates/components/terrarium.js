import * as THREE from 'three';
import { GUI } from 'three/addons/libs/lil-gui.module.min.js';

const container = document.getElementById("terrariumDiv");

const gui = new GUI( { container: container, title: 'Settings' } );
const renderer = new THREE.WebGLRenderer();
const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(75, container.clientWidth / container.clientHeight, 0.1, 1000);

const planetRadius = 15;
const geometry = new THREE.SphereGeometry( planetRadius, 32, 16 );
const material = new THREE.MeshPhongMaterial( { color: 0xCD853F } );
const planet = new THREE.Mesh( geometry, material );
const light = new THREE.DirectionalLight(0xffffff, 1);

function newBlob(color, radius, theta, y) {
	const geo = new THREE.SphereGeometry( radius, 32, 16 );
	const mat = new THREE.MeshPhongMaterial( { color: color } );
	const blob = new THREE.Mesh( geo, mat );
	const blobOrbit = new THREE.Spherical(planetRadius + radius - radius / 5, theta, y);
	blob.position.setFromSpherical(blobOrbit);

	planet.add( blob );
	return { blob, blobOrbit };
}

const { blob: blob1, blobOrbit: blobOrbit1 } = newBlob(0x0000ff, 0.5, THREE.MathUtils.degToRad(23), THREE.MathUtils.degToRad(23));
const { blob: blob2, blobOrbit: blobOrbit2 } = newBlob(0xff00ff, 0.5, THREE.MathUtils.degToRad(23), THREE.MathUtils.degToRad(50));

function animate() {
	// Speed of blob
	blobOrbit1.phi += 0.001;

	blob1.rotation.x += 0.01;
	blob1.rotation.y += 0.01;
	blob1.position.setFromSpherical(blobOrbit1);
	
	blobOrbit2.phi += 0.001;

	blob2.rotation.x += 0.01;
	blob2.rotation.y += 0.01;
	blob2.position.setFromSpherical(blobOrbit2);

	planet.rotation.x += 0.01;
	planet.rotation.y += 0.01;
	
	renderer.render(scene, camera);
}

camera.position.z = 30;

gui.domElement.style.zIndex = '10';
gui.domElement.style.position = 'absolute';
gui.add(camera.position, 'z', 20, 50).name("zoom");

light.position.set(0, 10, 0);
light.target.position.set(-5, 0, 0);

renderer.setSize(container.clientWidth, container.clientHeight);
renderer.setAnimationLoop(animate);

scene.background = new THREE.Color().setHex(0x001838);
scene.add( planet );
scene.add( light );
scene.add( light.target );

container.appendChild( renderer.domElement );

/* TODO */

// In the first place, each blob should have certain caracteristics.
// I am thinking of strenght, stamina, defence, notebook with informations that the blob gathered from the other blobs, bag of potions.
//
// Based on this caracteristics, when they meet an other blob, they can choose to do several action.
// They can choose to attack, cooperate, share data (like name, strenght, stamina or even data from an other blob)
// If both attack, they both deal damage based on the strenght
// If one attacks, whatever the other has chosen, he counters the attack with his defence and counter attack with half of his stamina
// If one dies, the other gets experience, increasing his stats based on how he won the fight. He also gets the potions from the other (so he can heal).
//
// After a certain timespan, the world ends.
// There is a ship at some place. If a blob gets to the ship, then it gets to escape and survive.
// There is no need to fight to get to the ship.
// If your blob escapes, you get to keep it with its newly trained characteristics and bag of potion or whatever it grabbed on the planet and you can send it in a new world.
